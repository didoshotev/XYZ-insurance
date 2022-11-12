import {
    Box,
    Button,
    Container,
    Flex,
    HStack,
    IconButton,
    useBreakpointValue,
    useColorModeValue,
} from '@chakra-ui/react'
import * as React from 'react'
import { FiMenu } from 'react-icons/fi'
import { Link } from 'react-router-dom'

export const Header = () => {
    const isDesktop = useBreakpointValue({
        base: false,
        lg: true,
    })
    return (
        <Box
            as="section"
            pb={{
                base: '12',
                md: '24',
            }}
        >
            <Box as="nav" bg="bg-surface" boxShadow={useColorModeValue('sm', 'sm-dark')}>
                <Container
                    py={{
                        base: '4',
                        lg: '5',
                    }}
                >
                    <HStack spacing="10" justify="space-between">
                        {isDesktop ? (
                            <Flex justify="space-between" flex="1">
                                <HStack spacing="3">
                                    <Button variant="ghost">
                                        <Link to={'/customers'}>Customers</Link>
                                    </Button>
                                    <Button variant="ghost">
                                        <Link to={'/insurances'}>Insurances</Link>
                                    </Button>
                                    <Button variant="ghost">
                                        <Link to={'/contracts'}>Contracts</Link>
                                    </Button>
                                </HStack>
                            </Flex>
                        ) : (
                            <IconButton
                                variant="ghost"
                                icon={<FiMenu fontSize="1.25rem" />}
                                aria-label="Open Menu"
                            />
                        )}
                    </HStack>
                </Container>
            </Box>
        </Box>
    )
}