SHELL = /bin/bash
MAKEFLAGS += --no-print-directory

args := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
$(eval $(args):dummy;@:)

dummy:	# fake target to avoid erreur manipuling $(args)

venv:
	python3 -m venv py3_env
	. py3_env/bin/activate && pip install --upgrade -r requirement.txt


run_day:
	@. py3_env/bin/activate && python ./day$(args)/day$(args).py

# run_all:
# 	@. py3_env/bin/activate && for file in `ls ./day*/day*.py`; do echo -e "\n>> $$file"; python $$file; done
run_all:
	@time for file in `ls ./day*/day*.py`; do echo -e "\n>> $$file"; python $$file; done