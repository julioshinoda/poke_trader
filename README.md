# POKE TRADER

API that calculate a trade between 2 trainers with maximum of 6 pokemons each. 

## Services

- **1. Calculate Trade**

Calculate if a trade between trainers is fair. 

endpoint: POST https://trade-pokemon.herokuapp.com/trade

headers: Content-Type: application/json

request body

```
{
    "first_trainer_list":[
        {
            "name":"squirtle"
        },
        {
            "name":"charizard"
        }
    ],
     "second_trainer_list":[
        {
            "name":"cubone"
        },
         {
            "name":"cubone"
        },
           {
            "name":"nidorina"
        }
    ]
}
```

**Response**

Response status: 
- 200 

Response payload:

``` 
{
  "fair_trade": false
}
```

The field fair_trade, when is true means that is a fair trade, otherwise the trade is unfair


- **2. Get Trade**


Get all made trade calculation

endpoint: GET https://trade-pokemon.herokuapp.com/trade

headers: Content-Type: application/json


**Response**

Response status: 
- 200 

Response payload:

```
[
  {
    "id": 25,
    "first_trainer_list": [
      {
        "base_experience": 63,
        "id": 7,
        "name": "squirtle"
      },
      {
        "base_experience": 240,
        "id": 6,
        "name": "charizard"
      }
    ],
    "second_trainer_list": [
      {
        "base_experience": 64,
        "id": 104,
        "name": "cubone"
      },
      {
        "base_experience": 64,
        "id": 104,
        "name": "cubone"
      },
      {
        "base_experience": 128,
        "id": 30,
        "name": "nidorina"
      }
    ],
    "created_at": "2021-01-15T00:00:00Z"
  }
]
```

## Makefile

Makefile file has the following commands

- **run**: That run application
- **test**: Running the unit tests


In this file has the environment variables for running application

