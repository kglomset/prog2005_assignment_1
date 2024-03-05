# Book Information Service

## PROG2005 - Cloud Technologies Assignment 1
___
This API is a RESTful web service that lets the user search for certain 
book data relating to one or more specific countries based on an ISO code 
input. It also gives the user the opportunity to check the status of the 
dependencies used in this project and how long the service has been active for, 
which are also REST web services. 

## Use
The service is deployed on Render at https://prog2005-assignment-1-yvtj.onrender.com.
In addition to the given URL you have to also provide the endpoints with the correct
parameters for the given endpoint you want to use. This is detailed for each endpoint below.
## Endpoints
```
/librarystats/v1/bookcount/
/librarystats/v1/readership/
/librarystats/v1/status/

```
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

## Dependencies
The following services has been used for this project. Documentation for
the given services are provided in the links below:
- Gutendex API
  - http://129.241.150.113:8000/
- Language2Countries API
  - http://129.241.150.113:3000/
- REST Countries API
  - http://129.241.150.113:8080/
