# BCC University

### ⚠️⚠️⚠️

```
Submissions from 2023 students will have much higher priority than submissions from 2022, SAP, or higher students.
Please take note of this before planning to attempt this freepass challenge.
```

## 💌 Invitation Letter

Embracing the digital era, we recognize the need for transparent and efficient information management in elections. We're embarking on a groundbreaking project to redefine the democratic process and we want you on board!

We aim to create an innovative election platform with features like candidate campaign forums, user participation, and administrative control. Your expertise is vital, and we're eager to hear your ideas.

Join us on this exciting journey to reshape the future of elections. Your contribution is key!

## **⭐** Minimum Viable Product (MVP)

As we have mentioned earlier, we need technology that can support BCC Election in the future. Please consider these features below:

- New user can register account to the system ✔️
- User can login to the system ✔️
- User can edit their profile account ✔️
- User can view the candidate's posts ✔️
- User can comment on candidate’s posts ✔️
- Users can view information about the candidates ✔️
- Users can cast their votes for candidates during the specified election period ✔️
- Admin can promote user to candidate ✔️
- Admin can view the candidate’s posts ✔️
- Admin can set the start and end dates for the election period ✔️
- Admin can delete the user/candidate ✔️
- Admin can delete the candidate's posts ✔️
- Admin can delete user comment ✔️
- Candidate can create, update, delete a post ✔️

## **[🌎](https://emojipedia.org/globe-showing-americas)** Service Implementation

```
GIVEN => I am a new user
WHEN  => I register to the system
THEN  => System will record and return the visitor's username

GIVEN => I am a user
WHEN  => I log in to the system
THEN  => System will authenticate and grant access based on user credentials

GIVEN => I am a user
WHEN  => I edit my profile account
THEN  => The system will update my account with the new details

GIVEN => I am a user
WHEN  => I view a candidate's post
THEN  => System will display the selected candidate's post along with its details

GIVEN => I am a user
WHEN  => I comment on a candidate’s post
THEN  => System will record my comment and return it under the candidate’s post

GIVEN => I am a user
WHEN  => I took an action to view candidate's posts
THEN  => System will show a candidate's post

GIVEN => I am a user
WHEN  => I cast my vote for a candidate during the specified election period
THEN  => System will register my vote for the selected candidate

GIVEN => I am an admin
WHEN  => I promote a user to a candidate
THEN  => System will update the user's status to candidate

GIVEN => I am an admin
WHEN => I view a candidate’s posts
THEN => System will display the posts created by the candidate

GIVEN => I am an admin
WHEN  => I set the start and end dates for the election period
THEN  => System will update the election period accordingly

GIVEN => I am an admin
WHEN  => I delete a user or candidate
THEN  => System will remove the user/candidate from the system

GIVEN => I am an admin
WHEN  => I delete a candidate’s post
THEN  => System will show a deletion status message and delete relevant post

GIVEN => I am an admin
WHEN => I delete a user comment
THEN => System will remove the user comment from the candidate's post

GIVEN => I am a candidate
WHEN  => I create a new post
THEN  => System will record and show the creation status message

GIVEN => I am a candidate
WHEN  => I update my post
THEN  => System will apply the changes and show an update status message

GIVEN => I am a candidate
WHEN  => I delete one of my posts
THEN  => System will show a deletion status message and delete relevant post
```

## **👪** Entities and Actors

We want to see your perspective about these problems. You can define various types of entities or actors. One thing for sure, there is no
true or false statement to define the entities. As long as the results are understandable, then go for it! 🚀

```
(Entities)
User: Represents individuals who interact with the system. They can have different roles such as regular users, candidates, or administrators. Attributes include:
- id
- nim (student ID)
- username
- password
- name
- major
- faculty
- status (user | candidate | admin)
- description

Post: Represents the content created by users, particularly candidates in this context. Attributes include:
- id
- user_id (referring to the creator)
- title
- content
- timestamp

Comment: Represents the comments made by users on posts. Attributes include:
- id
- user_id (referring to the commenter)
- post_id (referring to the post being commented on)
- content
- timestamp

Election: Represents the election event. Attributes include:
- id
- start_date
- end_date
- timestamp

Vote: Represents the votes cast by users in the election. Attributes include:
- id
- user_id (referring to the voter)
- election_id (referring to the election being voted in)
- candidate_id (referring to the candidate being voted for)
- timestamp

(Actors)
- User: Individuals who interact with the system. They can have different roles such as regular users, candidates, or administrators.
- Candidate: Users who are running for election. They can create posts to campaign and receive comments from other users.
- Administrator: Users with special privileges to manage the election system, such as creating or closing elections, managing users, and managing posts.
```

## **📘** References

You might be overwhelmed by these requirements. Don’t worry, here’s a list of some tools that you could use (it’s not required to use all of them nor any of them):

1. [Example Project](https://github.com/meong1234/fintech)
2. [Git](https://try.github.io/)
3. [Cheatsheets](https://devhints.io/)
4. [REST API](https://restfulapi.net/)
5. [Insomnia REST Client](https://insomnia.rest/)
6. [Test-Driven Development](https://www.freecodecamp.org/news/test-driven-development-what-it-is-and-what-it-is-not-41fa6bca02a2/)
7. [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
8. [GraphQL](https://graphql.org/)
9. [gRPC](https://grpc.io/)
10. [Docker Compose](https://docs.docker.com/compose/install/)

## **🔪** Accepted Weapons

> BEFORE CHOOSING YOUR LANGUAGE, PLEASE VISIT OUR [CONVENTION](CONVENTION.md) ON THIS PROJECT
>
> **Any code that did not follow the convention will be rejected!**
>
> 1. Golang (preferred)
> 2. Java (preferred)
> 3. NodeJS
> 4. PHP

You are welcome to use any libraries or frameworks, but we appreciate if you use the popular ones.

## **🎒** Tasks

```
The implementation of this project MUST be in the form of a REST, gRPC, or GraphQL API (choose AT LEAST one type).
```

1. Fork this repository
2. Follow the project convention
3. Finish all service implementations
4. Write the installation guide of your back-end service in the section below

## **🧪** API Installation

> Write how to run your service in local or development environment here. If you use Docker to serve your DBMS or your server, you will receive bonus points for your submission.

### Prerequisites
Before you begin, ensure you have the following installed on your machine. If you want to use **Docker** instead of manual installation, you can **skip these prerequisites**.

- Node.js
- npm
- MariaDB

### Running with Docker
These instructions will help you run the project locally using Docker with MariaDB.
1. Clone the repository or download manually:
```
git clone --branch yoga-raditya-nala https://github.com/yogarn/freepass-2024
```
2. Navigate to the project directory:
```
cd freepass-2024
```
3. Run Docker Compose to start the application:
```
docker-compose up -d
```
### Running without Docker (MariaDB)
If you prefer not to use Docker, you can run the application with a local MariaDB instance.
1. Make sure your MariaDB/MySQL server is configured correctly
2. Clone the repository or download manually:
```
git clone --branch yoga-raditya-nala https://github.com/yogarn/freepass-2024
```
3. Navigate to the project directory:
```
cd freepass-2024
```
4. Install dependencies:
```
npm install
```
5. Configure your credentials and database configuration in `config/database.js`.
6. Start the application:
```
npm start
```
### Access the aplication and documentations
You can **access** the application at http://localhost:3000/api/ and the **documentation** can be accessed in your web browser at http://localhost:3000/api-docs/.
### Stop the application
To stop the application and shut down the Docker containers, run (if using Docker):
```
docker-compose down
```
If you are not using Docker and using npm, you can simply press `ctrl + c` to kill the application.

## **[📞](https://emojipedia.org/telephone-receiver)** Contact

Have any questions? You can contact either [Vinncent](https://www.instagram.com/centwong_) or [Izra](https://www.instagram.com/izrarya).

## **🎁** Submission

Please follow the instructions on the [Contributing guide](CONTRIBUTING.md).

![cheers](https://gifsec.com/wp-content/uploads/2022/10/cheers-gif-1.gif)

> This is not the only way to join us.
>
> **But, this is the _one and only way_ to instantly pass.**
