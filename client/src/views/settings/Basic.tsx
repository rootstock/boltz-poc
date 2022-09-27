// ** React Imports
import { useState, ChangeEvent } from 'react'

// ** MUI Imports
import Grid from '@mui/material/Grid'
import Select from '@mui/material/Select'
import MenuItem from '@mui/material/MenuItem'
import InputLabel from '@mui/material/InputLabel'
import CardContent from '@mui/material/CardContent'
import FormControl from '@mui/material/FormControl'
import Button from '@mui/material/Button'
import BoltzPocService from 'src/@core/utils/service'
import { useEffect } from 'react';


const BasicSettings = () => {
  // ** State
  const [data, setData] = useState<string>("mint")


  const onClick = async () => {
    const result = await BoltzPocService.SaveConfig({
      key: "swaptype",
      value: data
    });
    console.log('Config saved', result);
  }

  useEffect(()=>{
    BoltzPocService.GetConfig("swaptype").then((value)=>{
      console.log("current config", value)
      setData(value.Value)
    });
  },[])

  return (
    <CardContent>
      <form>
        <Grid container spacing={7}>

          <Grid item xs={12} sm={6}>
            <FormControl fullWidth>
              <InputLabel id='Role'>Payment Type</InputLabel>
              <Select label='Payment Type' defaultValue='mint' value={data} onChange={(event)=>setData(event.target.value)}>
                <MenuItem value='mint'>DOC by Mint</MenuItem>
                <MenuItem value='liquidity'>DOC by Liquidity</MenuItem>
                <MenuItem value='rbtc'>RBTC</MenuItem>
              </Select>
            </FormControl>
          </Grid>

          <Grid item xs={12}>
            <Button variant='contained' sx={{ marginRight: 3.5 }} onClick={()=> onClick()}>
              Save 
            </Button>
          </Grid>
        </Grid>
      </form>
    </CardContent>
  )
}

export default BasicSettings
