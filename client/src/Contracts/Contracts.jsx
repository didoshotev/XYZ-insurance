import { Box, Heading } from "@chakra-ui/react"
import axios from "axios"
import { useEffect } from "react"
import { useState } from "react"
import { Header } from "../components/Header/Header"

export const Contracts = () => {
    const [contracts, setContracts] = useState([])

    useEffect(() => {
        axios({
            method: 'get',
            url: 'http://localhost:5000/contracts',
        }).then(res => {
            setContracts(res.data)
        }).catch(err => {
            console.log('error while loading contracts');
            console.error(err);
        })
    }, [])

    return (
        <>
            <Header />
            {contracts.map((contract, index) => {
                return (
                    <Box key={contract.contractId}>
                        <Box>
                            {index + 1}.
                            <span> Details: {contract.details} |</span>
                            <span> Duration: {contract.duration} |</span>
                            <span> Final Price: {contract.finalPrice} |</span>
                            {/* <span> Sign up Date: {new Date(contract.signUpDate).toLocaleDateString()} |</span> */}
                            {/* <span> Validity Date: {new Date(contract.validityDate).toLocaleDateString()} |</span> */}
                        </Box>
                    </Box>
                )
            })}
        </>
    )
}