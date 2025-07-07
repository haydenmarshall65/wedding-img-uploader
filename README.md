# Wedding Image Uploader

This is a fun side project I began creating with the concept in mind of creating a way for people to upload pictures from a wedding, but only the bride/groom could get access to it in order to be able to download the photos.

## The Plan:

### 1. Create an Auth system
Rather than use an existing auth, I wanted to learn how to use tokens and cookies in order to keep people logged in. It starts with the middleware. Users should be able to access the API, register, and login and view the frontend without their tokens invalidating for 30 days.

#### 1. apiMiddleware.go

#### 2. corsMiddleware.go

#### 3. frontEndMiddleware.go

## 2. Create a backend

In this app, users should be able to log in and then use the image uploader. There should be end points for taking in photos and storing them in an S3 bucket (or just a local file for testing).

In this app, the owners (Bride/Groom) need to be able to download those images. The API should check to make sure the user that is downloading the image is authenticated AND the user is the owner of the app. Otherwise, it should not allow them to join.

The owners should be able to delete any images they choose.

## 4. Create a frontend

In this app, users should be able to view the login/register pages and use their username/password to login.

There should be an easy way to set up an owner account.

There should be a way for the owners to view and download the photos added by the users.