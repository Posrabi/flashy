all:
	pip3 install -r codegen/requirements.txt && python3 codegen/codegen.py
	bash build_protos.sh
	bash build.sh
