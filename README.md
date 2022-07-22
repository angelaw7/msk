# msk-mongo

Code for uploading sample data from JSON file to MongoDB database and publishing it via a NATS server

## Setup

1. On MongoDB, create a database named `msk` with a collection called `testing` (or anything to match `collectionName` in `main.go`) 
2. Move `fetchjson.json` to root directory, run the following script to create new `fetch_shorter` file with unique dummy IDs, then move `fetchjson.json` to /pub directory
```python
python change_dummy_ids.py
```
3. Create a `.env` file with variable `MONGODB_URI` (find this on MongoDB connection configuration) in the /pub directory

## Usage

1. Run the following commands (from the root directory) to start the NATS server on Docker
```
docker run --name nats --network nats --rm -p 4222:4222 nats
```
2. Run the following commands (from the root directory) to run the subscriber on Docker (flags can be used to change default args)
```
cd sub
docker run --rm --name sub --network="host"
```
3. Run the following commands to run the publisher to upload all sample data and then update database with versioning based on the `fetch_shorter.json` (or anything to match `fetchJSONFile` in `main.go`) file on Docker (flags can be used to change default args):
```
cd pub
docker run --rm --name pub --network="host" pub
```
