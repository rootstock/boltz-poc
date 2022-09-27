// ** MUI Imports
import Grid from '@mui/material/Grid'
import Typography from '@mui/material/Typography'

// ** Demo Components Imports
import NFT from 'src/views/cards/NFT'

const ResearchersBasic = () => {
  return (
    <Grid container spacing={6}>
      <Grid item xs={12} sx={{ paddingBottom: 4 }}>
        <Typography variant='h5'>NFTs</Typography>
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <NFT id={1}/>
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <NFT id={2} />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <NFT id={3} />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <NFT id={4} />
      </Grid>
      <Grid item xs={12} sm={6} md={4}>
        <NFT id={5} />
      </Grid>
    </Grid>
  )
}

export default ResearchersBasic
