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
  -pwd string
    	Mongo Password (Optional). If username is provided and password is not the user will be prompted for it
  -url string
    	Mongo URL
  -user string
    	Mongo Username (Optional)
```

### Sample
```sh
$ csv2mongo -file=./CodingProjects.csv -url=127.0.0.1 -db=myapps -coll=apps
```

### Sample - Mongo Authentication
```sh
$ csv2mongo -file=./CodingProjects.csv -url=127.0.0.1 -db=myapps -coll=apps -user=user1 -pwd=pwd1
```

### Sample - Mongo Authentication with Pwd Prompt
```sh
$ csv2mongo -file=./CodingProjects.csv -url=127.0.0.1 -db=myapps -coll=apps -user=user1
Please enter the password for user1: 
```