INSERT INTO `periphery` (`id`, `name`, `is_measurable`)
VALUES (1, 'liquidPump', 0),
       (2, 'humiditySensor', 1),
       (3, 'led', 0);

INSERT INTO `job` (`id`, `name`)
VALUES (1, 'ledBlink'),
       (2, 'measureHumidity'),
       (3, 'pumpLiquid');

INSERT INTO `raspberry` (`id`, `name`, `plan_id`, `address`, `user`, `password`, `is_active`)
VALUES (1, 'Alfons', 1, 'rasp', '', '', 1);

INSERT INTO `pinout` (`raspberry_id`, `periphery_id`, `pin`)
VALUES (1, 1, 1),
       (1, 2, 2),
       (1, 3, 3);

INSERT INTO `plan` (`id`, `name`)
VALUES (1, 'mungBeanSprouts');

INSERT INTO `plan_value` (`id`, `plan_id`, `name`)
VALUES (1, 1, 'liquidPumpInterval'),
       (2, 1, 'liquidPumpDuration');

INSERT INTO `raspberry_plan_value` (`raspberry_id`, `plan_value_id`, `value`, `note`)
VALUES (1, 1, 8, 'in hours'),
       (1, 2, 3, 'in seconds');

