testenv:
	cd postgres && docker compose up -d
reset_testenv:
	cd postgres && docker compose down
	docker volume rm postgres_pg_data || echo no-volume
	make testenv
stop_testenv:
	cd postgres && docker compose down




# docker run -p 27017:27017 -d --name mongo_testenv mongodb/mongodb-community-server:latest \
# 	|| docker restart mongo_testenv