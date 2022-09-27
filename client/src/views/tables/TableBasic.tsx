// ** MUI Imports
import Paper from '@mui/material/Paper'
import Table from '@mui/material/Table'
import TableRow from '@mui/material/TableRow'
import TableHead from '@mui/material/TableHead'
import TableBody from '@mui/material/TableBody'
import TableCell from '@mui/material/TableCell'
import TableContainer from '@mui/material/TableContainer'
import { useState } from 'react';
import { useEffect } from 'react';

const createData = (name: string, calories: number, fat: number, carbs: number, protein: number) => {
  return { name, calories, fat, carbs, protein }
}

const rows = [
  createData('Frozen yoghurt', 159, 6.0, 24, 4.0),
  createData('Ice cream sandwich', 237, 9.0, 37, 4.3),
  createData('Eclair', 262, 16.0, 24, 6.0),
  createData('Cupcake', 305, 3.7, 67, 4.3),
  createData('Gingerbread', 356, 16.0, 49, 3.9)
]

const TableBasic = () => {  
  const [data, setData] = useState<[any]>()
  const [isLoading, setLoading] = useState(false)

  useEffect(() => {
    setLoading(true)
    fetch('http://localhost:8080/payment/')
      .then((res) => res.json())
      .then((data) => {
        setData(data)
        setLoading(false)
      })
  }, [])
  return (
    <TableContainer component={Paper}>
      <Table sx={{ minWidth: 650 }} aria-label='simple table'>
        <TableHead>
          <TableRow>
            <TableCell>#</TableCell>
            <TableCell align='right'>SATs</TableCell>
            <TableCell align='right'>Hash</TableCell>
            <TableCell align='right'>Tx ID</TableCell>
            <TableCell align='right'>Status</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>

          {data ? (
              data.length ? (
                data.map((todo, i) => {
                  return <TableRow key={todo.PreimageHash}>
                      <TableCell component='th' scope='row'>{i} </TableCell>
                      <TableCell align='right'>{todo.Amount} </TableCell>
                      <TableCell align='right'>{todo.PreimageHash.substring(0,5)}</TableCell>
                      <TableCell align='right'>{todo.Tx}</TableCell>
                      <TableCell align='right'>{todo.Status}</TableCell>
                  </TableRow>
                })
              ) : (
                <TableRow>
                  <TableCell colSpan={4}>No payments yet.</TableCell> 
                </TableRow>
              )
            ) : (
              <TableRow>
                <TableCell colSpan={4}>No payments yet.</TableCell> 
              </TableRow>
            )}
        </TableBody>
      </Table>
    </TableContainer>
  )
}

export default TableBasic
