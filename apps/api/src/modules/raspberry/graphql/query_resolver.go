package graphql

import (
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
	peripheryGraphql "main/modules/periphery/graphql"
	"shared/model/entity/periphery"
	"shared/model/entity/raspberry"
	"time"
)

func NewRaspberryQueryResolver(db *gorm.DB) *RaspberryQueryResolver {
	return &RaspberryQueryResolver{db}
}

type RaspberryQueryResolver struct {
	db *gorm.DB
}

type RaspberryListArgs struct {
	FromTime graphql.Time
	ToTime   graphql.Time
}

// todo how will whole graphql look? refactor this...
func (qr *RaspberryQueryResolver) RaspberryList(args *RaspberryListArgs) ([]*raspberryResolver, error) {
	var raspberries []*raspberry.Entity
	qr.db.Find(&raspberries)

	var raspIds []uint
	for _, rasp := range raspberries {
		raspIds = append(raspIds, rasp.ID)
	}

	var peripheries []struct {
		Id           uint
		Name         string
		RaspberryId  uint
		IsMeasurable bool
	}
	qr.db.
		Raw(`
	SELECT periphery.id, periphery.name, periphery.is_measurable, pinout.raspberry_id
	FROM periphery
	INNER JOIN pinout ON pinout.periphery_id = periphery.id AND pinout.raspberry_id IN (?)
`, raspIds).Scan(&peripheries)

	raspIdToPeripheries := make(map[uint][]*peripheryGraphql.PeripheryResolver)
	for _, peri := range peripheries {
		raspIdToPeripheries[peri.RaspberryId] = append(raspIdToPeripheries[peri.RaspberryId], &peripheryGraphql.PeripheryResolver{
			Entity: &periphery.Entity{
				ID:           peri.Id,
				Name:         peri.Name,
				IsMeasurable: peri.IsMeasurable,
			},
			Values: make([]*peripheryGraphql.PeripheryValueResolver, 0),
		})
	}

	var raspResolvers []*raspberryResolver
	for _, rasp := range raspberries {
		raspResolvers = append(raspResolvers, &raspberryResolver{rasp, raspIdToPeripheries[rasp.ID]})
	}

	var values []*struct {
		RaspberryId uint
		PeripheryId uint
		Value       float64
		DateTime    time.Time
	}
	qr.db.Raw(`
	SELECT AVG(m.value) AS value, m.created_at AS date_time, w.raspberry_id, m.periphery_id
	FROM measurement AS m
	INNER JOIN work AS w ON m.work_id = w.id AND w.raspberry_id IN (?)
	WHERE m.created_at BETWEEN ? AND ?
	GROUP BY w.raspberry_id, m.periphery_id, YEAR(m.created_at), MONTH(m.created_at), DAY(m.created_at), HOUR(m.created_at)
	ORDER BY m.created_at ASC
`, raspIds, args.FromTime.Time, args.ToTime.Time).Scan(&values)

	valueMap := make(map[uint]map[uint][]*peripheryGraphql.PeripheryValueResolver)

	for _, val := range values {
		if _, ok := valueMap[val.RaspberryId]; !ok {
			valueMap[val.RaspberryId] = make(map[uint][]*peripheryGraphql.PeripheryValueResolver)
		}
		t := val.DateTime
		valueMap[val.RaspberryId][val.PeripheryId] = append(valueMap[val.RaspberryId][val.PeripheryId], &peripheryGraphql.PeripheryValueResolver{
			Value_:    val.Value,
			DateTime_: time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0, t.Location()),
		})
	}

	for raspId, peripheries := range raspIdToPeripheries {
		for _, peri := range peripheries {
			if list, ok := valueMap[raspId][peri.Entity.ID]; ok {
				peri.Values = list
			}
		}
	}

	return raspResolvers, nil
}
