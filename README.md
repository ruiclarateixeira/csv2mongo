# CSV2MONGO

Simple command line utilty to import a csv file into a mongodb collection.

## Usage

```sh
$ csv2mongo -h
Usage of csv2mongo:
  -coll string
    	Collection Name
  -db string
    	Target Mongo Db Name
  -file string
    	Input File Path
  -url string
    	Mongo URL
```

### Sample

```sh
csv2mongo -file=./CodingProjects.csv -url=127.0.0.1 -db=myapps -coll=apps
```