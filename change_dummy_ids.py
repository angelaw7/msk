import json


def change_dummy_ids():
    """
    Short script for changing the dmp_sample_id for each sample to unique strings
    """

    filename = 'fetchjson.json'
    new_filename = 'fetch_shorter.json'

    with open(filename) as f:
        data = json.load(f)

    results = data["results"]

    for i in range(len(results)):
        result = results[i]
        result["meta-data"]["dmp_sample_id"] = str(i)

    with open(new_filename, 'w') as f:
        json.dump(data, f)


if __name__ == "__main__":
    change_dummy_ids()
