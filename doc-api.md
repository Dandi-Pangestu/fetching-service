**Register**
----
  Register new user to the system.

* **URL**

  http://localhost:4000/api/register

* **Method:**

  `POST`
  
*  **Headers**

   None
  
*  **URL Params**

   None

* **Data Params**

  **Required:**
   
  `{ "name" : "John", "phone": "08686767", "role": "admin" }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ "token" : "jwttoken", "password": "randompassword" }`
 
* **Error Response:**

  * **Code:** 422 ERR VALIDATION <br />
    **Content:** `[{ "msg": "Invalid value", "param": "name", "location": "body" }]`

  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error." }`

**Login**
----
  Login to the system.

* **URL**

  http://localhost:4000/api/login

* **Method:**

  `POST`
  
*  **Headers**

   None
  
*  **URL Params**

   None

* **Data Params**

  **Required:**
   
  `{ "phone": "08686767", "password": "yourpassword" }`

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ "token" : "jwttoken" }`
 
* **Error Response:**

  * **Code:** 422 ERR VALIDATION <br />
    **Content:** `[{ "msg": "Invalid value", "param": "phone", "location": "body" }]`
    
  OR
  
  * **Code:** 401 UNAUTHENTICATED <br />
    **Content:** `{ "message": "Unauthenticated." }`

  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error." }`

**Profile**
----
  Return user data (profile).

* **URL**

  http://localhost:4000/api/profile

* **Method:**

  `GET`
  
*  **Headers**

   **Required:**
   
   `Authorization: Bearer jwttokenplacehere`
  
*  **URL Params**

   None

* **Data Params**

  None

* **Success Response:**

  * **Code:** 200 <br />
    **Content:** `{ "name" : "John", "phone": "08686767", "role": "admin" }`
 
* **Error Response:**

  * **Code:** 401 UNAUTHENTICATED <br />
    **Content:** `{ "message": "Unauthenticated." }`

  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error." }`