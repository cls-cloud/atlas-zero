SUBMODULES := system monitor resource

.PHONY: build-all
build-all:
	@for dir in $(SUBMODULES); do \
	  $(MAKE) -C $$dir build; \
	done

.PHONY: run
run:
	./bin/traefik/traefik --configfile=./bin/traefik/traefik.yaml