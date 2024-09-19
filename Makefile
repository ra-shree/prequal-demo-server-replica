.PHONY: run-containers
run-containers:
	podman run --rm -d -p 9001:1233 --name server1 docker.io/rashree2023/load-balancer-probe-replica:v2
	podman run --rm -d -p 9002:1233 --name server2 docker.io/rashree2023/load-balancer-probe-replica:v2
	podman run --rm -d -p 9003:1233 --name server3 docker.io/rashree2023/load-balancer-probe-replica:v2
