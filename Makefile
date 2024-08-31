testenv:
	docker run -p 27017:27017 -d --name mongo_testenv mongodb/mongodb-community-server:latest \
		|| docker restart mongo_testenv
