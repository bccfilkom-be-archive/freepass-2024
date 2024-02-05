# PEMILU BERSIH PROJECT

## Dependencies:

```
Submissions from 2023 students will have much higher priority than submissions from 2022, SAP, or higher students.
Please take note of this before planning to attempt this freepass challenge.
```

## Build:


## Problem Statement:
Tingginya tingkat hoax dalam periode pemilu ini menjadikan saya membangun aplikasi Pemilu Bersih yang membantu masyarakt memperoleh informasi secara benar, masyarakat dapat mengetahui gagasan dari setiap paslon tanpa taakut apakah hal tersebut hoax atau tidak karena hanya paslon saja yang bisa menulis gagasan2 mereka.

## **⭐** Minimum Viable Product (MVP)

content management system:
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

Meet the actor :
-user
-candidate
-admin

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

```
(entity) users
- id
- name
- email
- phone_number
- role
- status
- timestamp

(entity) post
-id
-title
-description
-status
-timestamp

(entity) comment
-id
-user_id
-post_id
-comment
-status
-timestamp
```

