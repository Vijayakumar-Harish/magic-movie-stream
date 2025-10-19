import React from "react";
import { useState, useEffect } from "react";
import Container from "react-bootstrap/esm/Container";
import Button from "react-bootstrap/esm/Button";
import Form from "react-bootstrap/Form";
import axiosConfig from "../../api/axiosConfig";
import { useNavigate, Link } from "react-router-dom";

const Register = () => {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [favouriteGenres, setFavouriteGenres] = useState([]);
  const [genres, setGenres] = useState([]);

  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();
  return <h2>Register</h2>;
};

export default Register;
