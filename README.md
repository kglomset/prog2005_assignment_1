# Book Information Service

## PROG2005 - Cloud Technologies Assignment 1
___
This API is a RESTful web service that lets the user search for certain 
book data relating to one or more specific countries based on an ISO code 
input. It also gives the user the opportunity to check the status of the 
dependencies used in this project and how long the service has been active for, 
which are also REST web services. 
This web service contains three different end points detailed below.

## Endpoints
___
### Bookcount
```/librarystats/v1/bookcount/```

Example requests:

bookcount/?language=no

bookcount/?language=no,fi

Example response:


```[
{
"language": "no",
"books": 21,
"authors": 14,
"fraction": 0.0005
},
{
"language": "fi",
"books": 2798,
"authors": 228,
"fraction": 0.0671
}
]
```

### Readership
```/librarystats/v1/readership/```

Example requests:
- readership/no
- readership/sv/?limit=2

Example response:
```[
{
"country": "Norway",
"isocode": "NO",
"books": 21,
"authors": 14,
"readership": 5379475
},
{
"country": "Svalbard and Jan Mayen",
"isocode": "SJ",
"books": 21,
"authors": 14,
"readership": 2562
},
{
"country": "Iceland",
"isocode": "IS",
"books": 21,
"authors": 14,
"readership": 366425
}
]
```

### Status service
/librarystats/v1/status/

Response:
```
{
"gutendexapi": "<http status code for gutendex API>",
"languageapi: "<http status code for language2countries API>", 
"countriesapi": "<http status code for restcountries API>",
"version": "v1",
"uptime": <time in seconds from the last service restart>
}

```

## Deployment
Render
## Dependencies
