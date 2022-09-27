// ** MUI Imports
import Grid from '@mui/material/Grid'
import Typography from '@mui/material/Typography'

// ** Demo Components Imports
import Researcher from 'src/views/cards/Researcher'

const ResearchersBasic = () => {
  return (
    <Grid container spacing={6}>
      <Grid item xs={12} sx={{ paddingBottom: 4 }}>
        <Typography variant='h5'>Researchers</Typography>
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Shreemoy"/>
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Pato" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Raul" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Sergio" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Fede" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Juli" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Ramses" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Ilan" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Nico" />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <Researcher name="Javi" />
      </Grid>
    </Grid>
  )
}

export default ResearchersBasic
