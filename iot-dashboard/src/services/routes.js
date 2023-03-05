import axios from "axios";

const registerUser = async (email, password) => {
  return fetch("http://localhost:9090/auth/register", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
};

const signInUser = async (email, password) => {
  return fetch("http://localhost:9090/auth/login", {
    method: "POST",
    body: JSON.stringify({ email, password }),
  });
};

const createProject = async (
  name,
  dbUrl,
  dbAuthKey,
  bucketName,
  dbProjectName
) => {
  const header = new Headers();
  const token = localStorage.getItem("token");
  header.append("Authorization", `Bearer ${token}`);

  return fetch("/project", {
    method: "POST",
    headers: header,
    body: JSON.stringify({ name, dbUrl, dbAuthKey, bucketName, dbProjectName }),
  });
};

// const fetchProjectInfo = async () => {
//     return fetch("http://localhost:9090/protected/project", {
//       method: "get",
//     //   body: JSON.stringify({}),
//     });
//   };

export { registerUser, signInUser, createProject };
