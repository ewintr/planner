
sync-run:
	cd sync && PLANNER_DB_PATH=test.db PLANNER_PORT=8092 PLANNER_API_KEY=testKey go run .

sync-build-and-push:
	cd sync && docker build . -t codeberg.org/ewintr/syncservice && docker push codeberg.org/ewintr/syncservice
