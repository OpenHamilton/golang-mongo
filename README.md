#consuming hsr gtfs data

##This program downloads the zip file for hamilton's gtfs data, extracts it and loads it into the mongo db.

###VERSIONING

This program was written using the latest version to date(Aug 2014) of each the following except for mongodb which is an old version but this code should work on the latest version of mongodb too.

- go version 1.3
- mongodb v2.0.6
- mgo v.2 http://labix.org/mgo
 - You must install the mgo mongodb driver for Go before you can import it into your Go projects.

###PROGRAMS FILES BREAKDOWN

* The gsrgtfs/ and  lib/ would be located in $GOHOME/src/{username}/ or however you installed Go.
* The lib/check.go file contains an error function which kills the program and must be instead handled if this program is ever taken to production.
* The lib/gtfsStructs.go file contains definitions for our objects corresponding to the lines of data in each of the respective files in the zip archive. 
* The lib/Downloadzipfile.go file contains the script for downloading and saving the zip file.
* The lib/Loadmgo.go file contains the script to extract the zip file and hand off each files to the functions that follow:
 * The lib/loadmgoroutes.go file loads the routes.txt file into the mongodb.
 * The lib/loadmgostops.go file loads the stops.txt file into the mongodb.
 * The lib/loadmgostoptimes.go file loads the stopstimes.txt file into the mongodb.
 * The lib/loadmgotrips.go file loads the trips.txt file into the mongodb.

LICENSE: Creative Commons
Attribution 4.0 International
