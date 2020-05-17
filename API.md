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
    
**Fetching**
----
  Return response from resource and add usd price.

* **URL**

  http://localhost:8080/api/fetching/fetch

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
    **Content:** `[{"uuid": "8a23fcab-ef67-48b8-8ba1-7055ea91ea3b", "area_kota": "PANDEGLANG", "area_provinsi": "BANTEN", "komoditas": "Ikan Tuna Test Update", "price": "80000", "price_in_usd": "5.376922", "size": "80", "tgl_parsed": "2019-11-11T17:00:00.000Z", "timestamp": "1573491600"}]`
 
* **Error Response:**

  * **Code:** 401 UNAUTHENTICATED <br />
    **Content:** `{ "message": "Unauthenticated." }`

  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error." }`
    
**Aggregate**
----
  Return response from resource and aggregate it.

* **URL**

  http://localhost:8080/api/fetching/aggregate

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
    **Content:** `[{"area_provinsi": "BANTEN", "tahun": "2019", "minggu_ke": "46", "min": "45000", "max": "88000", "median": "4.000000", "avg": "61428.571429"}]`
 
* **Error Response:**

  * **Code:** 401 UNAUTHENTICATED <br />
    **Content:** `{ "message": "Unauthenticated." }`

  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error." }`
    
**Claim Token**
----
  Return jwt private claim.

* **URL**

  http://localhost:8080/api/fetching/claim-token

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
    **Content:** `{"name": "dandi", "phone": "1234", "role": "admin", "timestamp": "1589725816304"}`
 
* **Error Response:**

  * **Code:** 401 UNAUTHENTICATED <br />
    **Content:** `{ "message": "Unauthenticated." }`

  OR

  * **Code:** 500 INTERNAL SERVER ERR <br />
    **Content:** `{ "message": "Internal server error." }`