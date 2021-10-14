run: 
	make -j 2 backend frontend

build: 
	make -j 2 build-backend build-frontend

frontend: 
	make -C webapp/ run

backend:
	make -C server/ run

build-backend:
	make -C server build

build-frontend: 
	make -C webapp build
