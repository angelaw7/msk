# msk-mongo

Code for uploading sample data from JSON file to MongoDB database

## Setup

1. On MongoDB, create a database named `msk` with a collection called `sample_data`
2. Move a `edited_json.json` file with sample data to upload all data to the database (edited so that the `dmp_sample_id`s are all unique) 
3. Move a `new_input.json` file with new sample data into this directory
4. Create a `.env` file with variable `MONGODB_URI` (find this on MongoDB connection configuration)
5. Run the following command to upload all sample data and then update it based on the `new_input.json` file:
```go
go run main.go
```
