.DEFAULT_GOAL=no_target

no_target:
	@echo "Bro..."

install:
	@read -p "Enter SSH password: " PASS \
	&& ansible-playbook server-install.yaml -i honza@163.172.105.132, -v --extra-vars "ansible_sudo_pass=$${PASS}"
