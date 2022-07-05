# Simple REST based example

This is a simple deployment demonstrating a RESTful middleware interacting with MongoDB server.

This example is deployed in a [docker-compose network](../deployment/docker-compose.yaml).

To work with this example please use this [script](../scripts/compose.sh). Run the following commands:

* `./scripts/mongodb/ops.sh build` - Build containers
* `./scripts/mongodb/ops.sh start` - Start the RESTFul server
* `./scripts/mongodb/ops.sh stop` - Stop the RESTFul server
* `./scripts/mongodb/ops.sh clean` - Remove docker images and containers from your system.

When you have your REST server run, open a browser and target `http://localhost:4040` and this will return data extracted from the mongodb.