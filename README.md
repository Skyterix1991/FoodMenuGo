Default port that go is listening is 8080.


**Get all food**
----

* **URL**

  /food/

* **Method:**
  `GET`
  
* **Success Response:**
  
  * **Code:** 200 OK <br />
  * **Content:** `
    [
        {
            "Id": 1,
            "Name": "Cupcake",
            "Ingredients": [
                "Butter",
                "Sugar"
            ]
        },
        {
            "Id": 2,
            "Name": "Hotdog",
            "Ingredients": [
                "Sausage",
                "Roll"
            ]
        },
        {
            "Id": 3,
            "Name": "Cheeseburger",
            "Ingredients": [
                "Roll",
                "Meat",
                "Cheese"
            ]
        },
        {
            "Id": 0,
            "Name": "Salty mountain",
            "Ingredients": [
                "Salt",
                "More salt"
            ]
        }
    ]`
 
**Create food**
----

* **URL**

  /food/

* **Method:**
  `POST`
  
*  **JSON Body**
   `{
            "Id": 4,
            "Name": "Cupcake",
            "Ingredients": [
                "Butter",
                "Sugar"
            ]
        }`

* **Success Response:**
  * **Code:** 201 Created <br />
 
**Delete food**
----

* **URL**

  /food/{id}/

* **Method:**
  `DELETE`

* **Success Response:**
  * **Code:** 200 OK <br />
 
* **Error Response:**
  * **Code:** 404 Not found <br />

**Get food by id**
----

* **URL**

  /food/{id}/

* **Method:**
  `GET`

* **Success Response:**
  * **Code:** 200 OK <br />
  * **Content:** `{
                             "Id": 1,
                             "Name": "Cupcake",
                             "Ingredients": [
                                 "Butter",
                                 "Sugar"
                             ]
                         }`
 
* **Error Response:**
  * **Code:** 404 Not found <br />

**Update food**
----

* **URL**

  /food/{id}/

* **Method:**
  `PATCH`

*  **JSON Body**
   `{
            "Id": 1,
            "Name": "Cupcake",
            "Ingredients": [
                "Butter",
                "Sugar"
            ]
        }`

* **Success Response:**
  * **Code:** 200 OK <br />
 
* **Error Response:**
  * **Code:** 404 Not found <br />
  
**Replace food**
  ----
  
  * **URL**
  
    /food/{id}/
  
  * **Method:**
    `PUT`
  
  *  **JSON Body**
     `{
              "Id": 1,
              "Name": "Cupcake",
              "Ingredients": [
                  "Butter",
                  "Sugar"
              ]
          }`
  
  * **Success Response:**
    * **Code:** 200 OK <br />
   
  * **Error Response:**
    * **Code:** 404 Not found <br />