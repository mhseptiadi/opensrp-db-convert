# opensrp-db-conver

There are 2 main files here. 
- generate/generate.go | it is used for generate the schema at the begining of the conversion
- convert.go | it is used to incrementaly convert the data

## How to Run
You can either build the code first or just run the code.
`go run generate/generate.go`

After that, you can start populate the conversion table using
`go run convert.go`

Add a crontab to populate the different periodically