# Simple CRUD using Echo and pgsql

## Routes API

### Obtener Reviews

    ```
    Obtiene todas las reviews de la base de datos
    Ruta: {GET}
        $URL/getReviews
    ```
    
### Obtener Review específica

    ```
    Obtiene una review a partir de su id
    Ruta: {GET}
        $URL/getReview/:id
    ```

### Crear una Review

    ```
    Crea una review. Necesita:
        - id del usuario {user_id} // NOT NULL
        - id del teclado {keyboard_id} // NOT NULL
        - descripción {description} // default ''
        - stars {stars} // default 0

    EJ:
        {
            "keyboard_id": 2,
            "user_id": 1,
            "description": "este teclado ta bonito",
            "stars": 4
        }
    Ruta: {POST}
        $URL/createReview
    ```

### Editar una Review

    ```
    Edita una review. Necesita:
        - id de la review {id} // NOT NULL
        - descripción {description} // Enviar valor previo, por cualquier problema
        - stars {stars} // Enviar valor previo, por cualquier problema

    EJ:
        {
            "id": 1,
            "stars": 3,
            "description": "creo q ya no es tan sex0"
        }
    Ruta: {PUT}
        $URL/editReview
    ```

### Eliminar una Review

    ```
    Elimina una review. Necesita:
        - id de la review {id} // NOT NULL

    EJ:
        {
            "id": 1,
        }
    Ruta: {DELETE}
        $URL/deleteReview
    ```