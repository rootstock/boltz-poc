// ** MUI Imports
import Box from '@mui/material/Box'
import Card from '@mui/material/Card'
import Button from '@mui/material/Button'
import Avatar from '@mui/material/Avatar'
import CardMedia from '@mui/material/CardMedia'
import Typography from '@mui/material/Typography'
import CardContent from '@mui/material/CardContent'
import AvatarGroup from '@mui/material/AvatarGroup'
import { GetInfoResponse, requestProvider, WebLNProvider } from 'webln';
import { useState } from 'react';
import { useEffect } from 'react';
import BoltzPocService from 'src/@core/utils/service'

declare global {
  interface Window {
    webln: WebLNProvider
  }
}

const Researcher = (props) => {

  const [lnwallet, setLnwallet] = useState<GetInfoResponse>()
  const [isLoading, setLoading] = useState(false)
  const [isEnabled, setEnabled] = useState(false)

  const isProviderEnabled = async () => {
    setTimeout(async () => {
      try {
        await requestProvider()
        setLnwallet(await window.webln.getInfo())
        setEnabled(true);
        console.log('webln enabled', lnwallet )
      } catch (e) {
        setEnabled(false);
        console.error('webln NOT enabled', e)
      }
      setLoading(false)
    }, 100)
  }

  useEffect(() => {
    setLoading(true)
    isProviderEnabled()
  }, [])

  if (isLoading) return <div>Loading...</div>


  const onClick = async (e: any) => {
    e.preventDefault()
    if (isEnabled) {
      try {
        const lninvoice = await BoltzPocService.CreatePayment()
        await window.webln.sendPayment(lninvoice);
      } catch (error) {
        // ignore
        console.error(error);
      }
    } else {
      //show QR code.
      //linkFallback(this.paymentRequest);
    }
  }
  return (
    <Card sx={{ position: 'relative' }}>
      <CardMedia sx={{ height: '12.625rem' }} image='/images/cards/background-user.png' />
      <Avatar
        alt='Robert Meyer'
        src={'/images/cards/' + props.name + '.jpg'}
        sx={{
          width: 75,
          height: 75,
          left: '1.313rem',
          top: '10.28125rem',
          position: 'absolute',
          border: theme => `0.25rem solid ${theme.palette.common.white}`
        }}
      />
      <CardContent>
        <Box
          sx={{
            mt: 5.75,
            mb: 8.75,
            display: 'flex',
            flexWrap: 'wrap',
            alignItems: 'center',
            justifyContent: 'space-between'
          }}
        >
          <Box sx={{ mr: 2, mb: 1, display: 'flex', flexDirection: 'column' }}>
            <Typography variant='h6'>{props.name}</Typography>
            <Typography variant='caption'>London, UK</Typography>
          </Box>
          <Button variant='contained' onClick={onClick}>Book</Button>
        </Box>
        <Box sx={{ gap: 2, display: 'flex', flexWrap: 'wrap', justifyContent: 'space-between', alignItems: 'center' }}>
          <Typography variant='subtitle2' sx={{ whiteSpace: 'nowrap', color: 'text.primary' }}>
            18 mutual friends
          </Typography>
          <AvatarGroup max={4}>
            <Avatar src='/images/avatars/8.png' alt='Alice Cobb' />
            <Avatar src='/images/avatars/7.png' alt='Jeffery Warner' />
            <Avatar src='/images/avatars/3.png' alt='Howard Lloyd' />
            <Avatar src='/images/avatars/2.png' alt='Bettie Dunn' />
            <Avatar src='/images/avatars/4.png' alt='Olivia Sparks' />
            <Avatar src='/images/avatars/5.png' alt='Jimmy Hanson' />
            <Avatar src='/images/avatars/6.png' alt='Hallie Richards' />
          </AvatarGroup>
        </Box>
      </CardContent>
    </Card>
  )
}

export default Researcher
