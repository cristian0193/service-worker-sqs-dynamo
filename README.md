# Service Worker SQS - Dynamo

## Tabla de contenido
1. [Contexto](#contexto)
2. [Tecnologías](#tecnologías)
3. [Arquitectura](#arquitectura)
    * [Estructura del proyecto](#estructura-del-proyecto)
5. [Despliegues](#despliegues)
    * [Local](#local)
6. [Endpoints](#endpoints)
7. [Queues](#queues)


<a name="contexto"></a>
# Contexto 📋

El objetivo de este template es proporcionar una estructura básica para crear un servicio de trabajador (worker) que establezca conexión con servicios de AWS, como SQS y RDS. Esta plantilla tiene como finalidad simplificar la construcción de servicios similares, al proporcionar una configuración sencilla y reducir el tiempo necesario para su desarrollo.

- **Estructurar un servicio worker**: Crear una estructura clara y organizada para desarrollar un servicio de trabajador que cumpla con los requisitos del proyecto.
- **Establecer conexión con servicios AWS**: Configurar la conexión con servicios de AWS, como Amazon SQS (Simple Queue Service) y RDS (Relational Database Service).
- **Simplificar la configuración**: Proporcionar una configuración sencilla que permita a los desarrolladores personalizar y adaptar el servicio según sus necesidades específicas.
- **Reducir tiempos de construcción**: Agilizar el proceso de desarrollo al brindar una plantilla predefinida y una estructura base que sirva como punto de partida para la construcción del servicio.
- **Mejorar la reutilización de código**: Fomentar la reutilización de código al proporcionar una estructura común y componentes predefinidos que puedan ser utilizados en diferentes proyectos de servicios de trabajador.


<a name="tecnologías"></a>
# Tecnologias 💻

**Dependencies** 🤝
Las siguientes dependencias se utilizan en el desarrollo para llevar a cabo depliegue de servidor http, conexiones SQS y RDS, entre otros.

* [github.com/aws/aws-sdk-go](https://github.com/aws/aws-sdk-go): SDK oficial de AWS para el lenguaje de programación Go.

* [github.com/labstack/echo/v4](https://github.com/labstack/echo): Echo es un framework de aplicaciones web Go de código abierto, extensible y centrado en el rendimiento.

* [go.uber.org/zap](https://pkg.go.dev/go.uber.org/zap): Permite la configuracion de logs.

**Framework**

* [Echo](https://echo.labstack.com/)

**Servicios AWS**

* [SQS](https://aws.amazon.com/es/sqs/)

**Bases de datos**

* [Dynamo](https://aws.amazon.com/es/dynamodb/)

<a name="arquitectura"></a>
# Arquitectura 🏢

Para del proyecto se toma como base los principios de las arquitecturas limpias, utilizando en este caso gran parte del concepto de **arquitectura multicapas**, lo cual permite la independencia de frameworks, entidades externas y UI, por medio de capas con responsabilidad únicas que permite ser testeables mediante el uso de sus interfaces. Como parte de las buenas prácticas la solución cuenta en su gran mayoría con la aplicación de los principios SOLID, garantizando un código limpio, mantenible, reutilizable y escalable.

![service-worker-sqs-dynamodb](https://github.com/cristian0193/service-worker-sqs-dynamodb/assets/11803196/f066be29-3b5b-47b9-ad8b-24db04f05d52)

<a name="estructura-del-proyecto"></a>
### * **Estructura del proyecto** 🧱

- [x] `config/`: establece las configuraciones iniciales a los servicios
  - [ ] `cmd/`: administra los recursos de llamados al api
    - [ ] `builder/`: construye cada una de las instancias transversales
- [x] `core/`: establece la logica de negocio
  - [ ] `domain/`: administracion de los datos de manera transversal
  - [ ] `usecases/`: define los casos de uso utilizados por el handler
- [x] `dataproviders/`: contiene la implementacion de los clients externos
  - [ ] `awsdynamodb/`: define el cliente para aws dynamodb
  - [ ] `awssqs/`: define el cliente para aws sqs
  - [ ] `consumer/`: define la logica para obtener los mensajes desde el consumidor
  - [ ] `processor/`: define el inicio del proceso para la lectura de mensajes desde SQS
  - [ ] `repository/`: define las consultas, actualizacion o inserciones a la base de datos
  - [ ] `server/`: define la configuracion para correr el server http
  - [ ] `utils/`: define las funciones transversales
- [x] `entrypoints/`: administra los recursos de llamados al api
  - [ ] `controllers/`: define los handler

<a name="despliegues"></a>
# Despliegues 🚀

Para la fase de despliegue a nivel local, se utilizaron algunas herramientas que nos permite agilizar este proceso. Como fase se mostrara el paso a paso:

> **Nota:** Para el proceso se deben definir las variables de ambiente que nos permite establecer conexion a los diferentes servicios.

```
APPLICATION_ID=
SERVER_PORT=
LOG_LEVEL=INFO

AWS_ACCESS_KEY=
AWS_SECRET_KEY=
AWS_REGION=

AWS_SQS_URL=
AWS_SQS_MAX_MESSAGES=
AWS_SQS_VISIBILITY_TIMEOUT=

AWS_DYNAMODB_TABLE=
```

<a name="local"></a>
### * **Local** 

En el proceso local podemos utilizar despliegues de contenedores con dynamo.

    1. Creacion de DynamoDB en AWS
        - https://aws.amazon.com/es/dynamodb/

    2. Creacion de SQS en AWS
        - https://aws.amazon.com/es/sqs/

    3. Definir las variables de entorno definidas en (Despliegues)

    4. Ejecutar el comando 'go run main.go'

    5. Puerto :8080 run

<a name="endpoints"></a>
# Endpoints 🤖

- **GET**    http://localhost:8080/sqs/:id
```
curl --location --request GET 'http://localhost:8080/sqs/:id'
```

- **Response**
```
  {
    "id": "7a312c5a-e69e-4935-9b33-5dc33919a76f",
    "message": "Hola Mundo!!",
    "date": "2023-06-13T17:48:05-05:00"
  }
```

<a name="queues"></a>
# Queues 📨

- **URL**    https://sqs.us-east-1.amazonaws.com/XXXXXXXX/service-worker-sqs-dynamodb


- **Message**
```
    {
      "message": "Hello World"
    }
```

# Author 🧑‍💻
```
- Christian Alexis Rodriguez Castillo
- Sr Software Engineer - Mercadolibre
- cristian010193@gmail.com
```
