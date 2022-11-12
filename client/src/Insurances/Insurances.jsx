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
import axios from "axios"
import { useEffect } from "react"
import { useState } from "react"
import { Header } from "../components/Header/Header"

export const Insurances = () => {
    const [insurances, setInsurances] = useState([])

    useEffect(() => {
        axios({
            method: 'get',
            url: 'http://localhost:5000/insurances',
        }).then(res => {
            setInsurances(res.data)
        }).catch(err => {
            console.log('error while loading insurances');
            console.error(err);
        })
    }, [])

    const formik = useFormik({
        initialValues: {
            name: "",
            type: "",
            price: 0.00
        },
        onSubmit: (values) => {
            axios({
                method: 'post',
                url: 'http://localhost:5000/insurances',
                data: {...values, price: +values.price},
            }).then(res => {
                console.log('success: ', res)
            })
                .catch(err => {
                    console.log('Could not create an Insurance');
                    console.error('error: ', err);
                })
        }
    });

    return (
        <FormikProvider value={formik}>
            <Header />
            <Flex justify="center">
                <Box bg="white" p={6} rounded="md">
                    <Heading pb={10} size={'lg'}>Create Insurance</Heading>
                    <form onSubmit={formik.handleSubmit}>
                        <VStack spacing={4} align="flex-start">
                            <FormControl>
                                <FormLabel htmlFor="name">Insurance Name</FormLabel>
                                <Input
                                    id="name"
                                    name="name"
                                    type="text"
                                    variant="filled"
                                    onChange={formik.handleChange}
                                    value={formik.values.name}
                                />
                            </FormControl>
                            <FormControl>
                                <FormLabel htmlFor="type">Insurance Type</FormLabel>
                                <Input
                                    id="type"
                                    name="type"
                                    type="text"
                                    variant="filled"
                                    onChange={formik.handleChange}
                                    value={formik.values.type}
                                />
                            </FormControl>
                            <Field name='price'>
                                {({ field, form }) => (
                                    <FormControl id='price'>
                                        <FormLabel htmlFor='price'>Price</FormLabel>
                                        <NumberInput
                                            id='price'
                                            {...field}
                                            onChange={(val) =>
                                                form.setFieldValue(field.name, val)
                                            }
                                            value={formik.values.price}
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
                            <Button type="submit" colorScheme="purple" width="full">
                                Submit
                            </Button>
                        </VStack>
                    </form>
                </Box>
            </Flex>
            {insurances.map(insurance => {
                return (
                    <Box key={insurance.insuranceId}>
                        <Box>
                            <span>name: {insurance.name} |</span>
                            <span> type: {insurance.type} |</span>
                            <span> ( price: {insurance.price})</span>
                        </Box>
                    </Box>
                )
            })}
        </FormikProvider>
    )
}