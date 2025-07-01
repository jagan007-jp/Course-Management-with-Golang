import { useState } from "react";
import { useNavigate } from "react-router-dom";
import {
  Box,
  Input,
  Button,
  Heading,
  FormControl,
  FormLabel,
  Stack,
  Container,
  useToast,
} from "@chakra-ui/react";
import Navbar from "./navbar";

export default function AddCourse(){
    const [courseName, setCourseName] = useState('')
    const [courseLink, setCourseLink] = useState('')
    const [courseImage, setCourseImage] = useState('')
    const token = localStorage.getItem("token")
    const toast = useToast()
    const navigate = useNavigate()
    const addCourse = async()=>{
        try{
        const res = await fetch("http://localhost:8080/courses/add",{
            method:"POST",
            headers: {"Content-Type":"application/json", Authorization: `${token}`},
            body: JSON.stringify({title: courseName, image: courseImage, link: courseLink})
        })

        const data = await res.json()
        if(res.ok){
            toast({
                title: "Added to courses",
                description: data.message,
                status: "success",
                duration: 3000,
                isClosable: true
            })
            navigate("/home")
        }
        }catch(err){
        toast({
            title: "Server error",
            description: err.message,
            status: "error",
            duration: 3000,
            isClosable: true
        })
    }
    }

    return(
        <>
        <Navbar/>
        <Box bg="gray.50" minH="100vh" py={10}>
      <Container maxW="md" bg="white" p={8} borderRadius="lg" boxShadow="md">
        <Heading textAlign="center" mb={6} color="blue.700">
          Add Course
        </Heading>
        <form>
          <Stack spacing={4}>
            <FormControl isRequired>
              <FormLabel>Course Name</FormLabel>
              <Input
                type="text"
                maxLength={40}
                value={courseName}
                onChange={(e) => setCourseName(e.target.value)}
                placeholder="Enter the name of the course"
              />
            </FormControl>

            <FormControl isRequired>
              <FormLabel>Course Image link</FormLabel>
              <Input
                type="text"
                maxLength={300}
                value={courseImage}
                onChange={(e) => setCourseImage(e.target.value)}
                placeholder="Enter Image link"
              />
            </FormControl>

            <FormControl isRequired>
              <FormLabel>Course Link</FormLabel>
              <Input
                type="text"
                maxLength={150}
                value={courseLink}
                onChange={(e) => setCourseLink(e.target.value)}
                placeholder="Enter Course link"
              />
            </FormControl>

            <Button colorScheme="blue" onClick={addCourse}>
              Add to courses
            </Button>
          </Stack>
        </form>
      </Container>
    </Box>
    </>
    )
}