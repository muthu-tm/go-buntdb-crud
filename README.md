# go-BuntDB-CRUD

    Tiny Go-Rest service created using Gorilla MUX to perform basic CRUD operation on BuntDB

MUX - https://github.com/gorilla/mux

BuntDB - https://github.com/tidwall/buntdb

REST Endpoints:

    - POST      http://localhost:3000/emp
        Add new employee to our in-memory, key/value database

        Ex Body params for this POST Request,

        ```
        {
	        "name": "Muthu",
	        "contact": "muthu@gmail.com",
	        "salary": 100000,
	        "is_active": true
        }
        ```

    - GET       http://localhost:3000/emp
        Fetch all employee as a list from database

    - GET       http://localhost:3000/emp/{id}
        Fetch a particular employee from database

    - PUT       http://localhost:3000/emp/{id}
        Update an employee detail in database

        Ex Body params for this PUT Request,

        ```
        {
	        "name": "Muthu TM",
	        "contact": "muthu@gmail.com",
	        "salary": 100000,
	        "is_active": false
        }
        ```

    - DELETE    http://localhost:3000/emp/{id}
        Delete an employee from database
