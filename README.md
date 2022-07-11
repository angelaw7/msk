# msk-mongo

Code for uploading sample data from JSON file to MongoDB database

## Setup

1. On MongoDB, create a database named `msk` with a collection called `testing` (or anything to match `collectionName` in `main.go`) 
2. Move `fetchjson.json` to root directory and run the following script to create new `fetch_shorter` file with unique dummy IDs
```python
python change_dummy_ids.py
```
3. Create a `.env` file with variable `MONGODB_URI` (find this on MongoDB connection configuration)
4. Run the following command to upload all sample data and then update database with versioning based on the `new_input.json` (or anything to match `fetchJSONFile` in `main.go`) file:
```go
go run main.go
```
