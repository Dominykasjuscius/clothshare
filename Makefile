run: 
	make -j 2 backend frontend

frontend: 
	make -C webapp/ run

backend:
	make -C server/ run
