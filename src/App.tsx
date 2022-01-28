import React from "react";
import { Header } from "./Components/index";
import "./App.css";
import { Box, Heading, Container, Text, Button, Stack, Icon, useColorModeValue, createIcon } from "@chakra-ui/react";
import { TodoList } from "./Components/todolist/todolist";

export default function App() {
  return (
    <>
      <Container maxW={"3xl"}>
        <Stack as={Box} textAlign={"center"} spacing={{ base: 8, md: 14 }} py={{ base: 20, md: 36 }}>
          <Header />
          <TodoList
            todos={[
              { title: "Do dishes", isCompleted: true },
              { title: "Mow the lown", isCompleted: false },
            ]}
          />
        </Stack>
      </Container>
    </>
  );
}
