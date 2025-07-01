import { useState } from "react";
import { setUser, setError } from "../slice/userSlice";
import { useSelector, useDispatch } from "react-redux";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import {
  Box,
  Input,
  Button,
  Heading,
  Text,
  FormControl,
  FormLabel,
  Stack,
  Container,
  useToast,
  Link,
  Select
} from "@chakra-ui/react";

export default function Login() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("");
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const toast = useToast();
  const [error, setError] = useState('');


  const handleLogin = async (e) => {
    const user1 = username.trim().toUpperCase()
    const pass1 = btoa(btoa(btoa(password.trim())))
    let pass2 = ''
    for(let i=0; i<pass1.length; i++){
      pass2 += ''+pass1.charCodeAt(i).toString(16);
    }
    console.log(pass2)
    e.preventDefault();
    try {
      const res = await fetch("http://localhost:8080/users/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ username:user1, password:pass2, role }),
      });
      const data = await res.json();
      if (!res.ok) {
        setError(data.message);
        toast({
          title: "Login Failed",
          description: data.message,
          status: "error",
          duration: 3000,
          isClosable: true,
        });
      } else {
        dispatch(setUser({username:user1}));
        localStorage.setItem("username", user1);
        localStorage.setItem("role",role);
        localStorage.setItem("token",data.token);
        toast({
          title: "Login Successful",
          description: `Welcome ${user1}!`,
          status: "success",
          duration: 3000,
          isClosable: true,
        });
        navigate("/home");
      }
    } catch (err) {
      setError("Server Error");
      toast({
        title: "Server Error",
        description: "Please try again later.",
        status: "error",
        duration: 3000,
        isClosable: true,
      });
    }
  };

  return (
    <Box bg="gray.50" minH="100vh" py={10}>
      <Container maxW="md" bg="white" p={8} borderRadius="lg" boxShadow="md">
        <Heading textAlign="center" mb={6} color="blue.700">
          Login
        </Heading>
        <form onSubmit={handleLogin}>
          <Stack spacing={4}>
            <FormControl isRequired isInvalid={error && username === ""}>
              <FormLabel>Username</FormLabel>
              <Input
                type="text"
                maxLength={20}
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                placeholder="Enter Username"
              />
            </FormControl>

            <FormControl isRequired isInvalid={error && password === ""}>
              <FormLabel>Password</FormLabel>
              <Input
                type="password"
                maxLength={100}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="Enter Password"
              />
            </FormControl>

            <FormControl isRequired isInvalid={error && role === ""}>
            <FormLabel>Role</FormLabel>
              <Select
                placeholder="Select Role"
                value={role}
                onChange={(e) => setRole(e.target.value)}
              >
                <option value="admin">Admin</option>
                <option value="user">User</option>
              </Select>
            </FormControl>

            <Button colorScheme="blue" type="submit">
              Login
            </Button>
            <Text fontSize="sm" color="gray.600" textAlign="center">
              New here?{" "}
              <Link as={RouterLink} to="/register" color="blue.500">
                Register
              </Link>
            </Text>

            {error && (
              <p style={{
                textAlign: "center",
                color: "red"
              }}>
                {error}
              </p>
            )}
          </Stack>
        </form>
      </Container>
    </Box>
  );
}
