CONFIGPATH = ./config
CONFIGNAME = local

all : run

run :
	CONFIGPATH=../$(CONFIGPATH) CONFIGNAME=$(CONFIGNAME) make -C src