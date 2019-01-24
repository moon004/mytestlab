test:
	bin/maketest.sh
	
check_go_path:
	echo "gpath: $(GOPATH)"
	GOPATH="$(GOPATH)" bin/check_path testlab
.PHONY: check_go_path