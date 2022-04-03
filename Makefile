gen:
	pip3 install -r codegen/requirements.txt && python3 codegen/codegen.py
pb:
ifdef ts
	bash build_protos.sh ts
else
	bash build_protos.sh
endif
code:
ifdef ts
	bash build.sh ts
else
	bash build.sh
endif
