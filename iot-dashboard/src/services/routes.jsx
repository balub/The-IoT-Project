import axios from "axios";

const registerUser = async (email, password) => {
  return fetch("http://localhost:9090/auth/register", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
};

const signInUser = async (email, password) => {
  fetch("http://localhost:9090/auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  })
    .then((v) => v.json)
    .then((v) => {
      console.log(v);
    })
    .catch((err) => {
      console.log(err);
    });
};

const createProject = async (title, dbUrl) => {
  fetch("http://localhost:9090/protected/project", {
    method: "POST",
    body: JSON.stringify({ title, dbUrl }),
  })
    .then((v) => v.json)
    .then((v) => {
      console.log(v);
    })
    .catch((err) => {
      console.log(err);
    });
};

const fetchProjectInfo = async () => {
  fetch("http://localhost:9090/protected/project", {
    method: "GET",
    body: JSON.stringify({ title, dbUrl }),
  })
    .then((v) => v.json)
    .then((v) => {
      console.log(v);
    })
    .catch((err) => {
      console.log(err);
    });
};

export { registerUser, signInUser, createProject, fetchProjectInfo };
