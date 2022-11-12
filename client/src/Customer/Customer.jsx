import { Field, FormikProvider, useFormik } from "formik";
import {
    Box,
    Button,
    Flex,
    FormControl,
    FormLabel,
    Heading,
    Input,
    NumberDecrementStepper,
    NumberIncrementStepper,
    NumberInput,
    NumberInputField,
    NumberInputStepper,
    UnorderedList,
    VStack
} from "@chakra-ui/react";
import { useEffect } from "react";
import { useState } from "react";
import axios from "axios";
import { Header } from "../components/Header/Header";

export const Customer = () => { 
    const [customers, setCustomers] = useState([])

    useEffect(() => {
        axios({
            method: 'get',
            url: 'http://localhost:5000/customers',
            // withCredentials: true
        }).then(res => {
            setCustomers(res.data)
        }).catch(err => {
            console.log('error while loading customers');
            console.error(err);
        })
    }, [])

    const formik = useFormik({
        initialValues: {
            personName: "",
            familiyName: "",
            dob: "",
            rating: 1
        },
        onSubmit: (values) => {
            axios({
                method: 'post',
                url: 'http://localhost:5000/customers',
                data: values,
            }).then(res => {
                console.log('success: ', res)
            })
                .catch(err => {
                    console.log('Could not create a customer');
                    console.error('error: ', err);
                })
        }
    });
    return (
        <FormikProvider value={formik}>
            <Header />
            <Flex justify="center">
                <Box bg="white" p={6} rounded="md">
                    <Heading pb={10} size={'lg'}>Create Customer</Heading>
                    <form onSubmit={formik.handleSubmit}>
                        <VStack spacing={4} align="flex-start">
                            <FormControl>
                                <FormLabel htmlFor="email">Person Name</FormLabel>
                                <Input
                                    id="personName"
                                    name="personName"
                                    type="text"
                                    variant="filled"
                                    onChange={formik.handleChange}
                                    value={formik.values.personName}
                                />
                            </FormControl>
                            <FormControl>
                                <FormLabel htmlFor="familiyName">Familiy Name</FormLabel>
                                <Input
                                    id="familiyName"
                                    name="familiyName"
                                    type="text"
                                    variant="filled"
                                    onChange={formik.handleChange}
                                    value={formik.values.familiyName}
                                />
                            </FormControl>

                            <Field name='rating'>
                                {({ field, form }) => (
                                    <FormControl id='rating'>
                                        <FormLabel htmlFor='rating'>Rating</FormLabel>
                                        <NumberInput
                                            id='rating'
                                            {...field}
                                            onChange={(val) =>
                                                form.setFieldValue(field.name, val)
                                            }
                                            value={formik.values.rating}
                                            min={0}
                                            max={10}
                                        >
                                            <NumberInputField />
                                            <NumberInputStepper>
                                                <NumberIncrementStepper />
                                                <NumberDecrementStepper />
                                            </NumberInputStepper>
                                        </NumberInput>
                                    </FormControl>
                                )}
                            </Field>

                            <FormControl>
                                <FormLabel htmlFor="dob">Date Of Birth</FormLabel>
                                <input
                                    id="dob"
                                    name="dob"
                                    type="date"
                                    onChange={formik.handleChange}
                                    value={formik.values.dob}
                                />
                            </FormControl>
                            <Button type="submit" colorScheme="purple" width="full">
                                Submit
                            </Button>
                        </VStack>
                    </form>
                </Box>
            </Flex>
            <UnorderedList>
                {customers.map(customer => {
                    return (
                        <Box key={customer.customerId}>
                            <Box>Name:
                                <span> {customer.personName}</span>
                                <span> {customer.familyName}</span>
                                <span> (rating: {customer.rating})</span>
                            </Box>
                        </Box>
                    )
                })}
            </UnorderedList>
        </FormikProvider>
    );

    return ( 
        <h1>Welcome to the Customer Page</h1>
    )
}