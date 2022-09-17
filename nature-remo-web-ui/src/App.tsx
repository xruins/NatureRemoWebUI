import React, { useEffect, useState } from 'react';
import { styled } from '@mui/material/styles';
import { AppBar, Box, CssBaseline, IconButton, Paper, Toolbar, Typography } from '@mui/material';
import './App.css';
import { AppliancesPanel } from './appliance';
import { getAppliances } from './remo/client';
import { Appliance } from './remo/model';
import { Refresh } from '@mui/icons-material';
import toast, { Toaster } from 'react-hot-toast';

const DivRoot = styled('div')(({ theme }) => ({
  flexGrow: 1,
  backgroundColor: theme.palette.background.default
}));


const TypographyTitle = styled(Typography)(({ theme }) => ({
  margin: theme.spacing(2),
  padding: theme.spacing(2)
}));

const PaperContent = styled(Paper)(({ theme }) => ({
  margin: theme.spacing(2),
  padding: theme.spacing(2)
}));

export default function App() {
  const [appliances, setAppliances] = React.useState<Appliance[]>([]);

  const handleApplianceButtonClick = () => {
    toast.promise(
      getAppliances(true),
      {
        loading: 'Updating appliances...',
        success: (response: Appliance[]) => {
          setAppliances(response)
          return `Succeeded to update appliances.`;
        },
        error: (err) => {
          let msg = "";
          msg += `Failed to update appliances \`${name}\``;
          msg += err !== undefined ? `\nDetail: ${err}` : '';
          msg += err.message !== undefined ? `\nMessage: ${err.message}` : '';
          return msg;
        },
      }
    )
  };

  React.useEffect(() => {
    document.title = "Remo WebUI";
    getAppliances(false).then(
      (appliances: Appliance[]) => {
        setAppliances(appliances);
      }
    ).catch(
      (err) => {
        let msg = "";
        msg += `Failed to get appliances \`${name}\``;
        msg += err !== undefined ? `\nDetail: ${err}` : '';
        msg += err.message !== undefined ? `\nMessage: ${err.message}` : '';
        toast.error(msg);
      },
    );
  }, []);

  return (
    <DivRoot>
      <Box sx={{ display: 'flex' }}>
        <CssBaseline />
        <AppBar position="static">
          <Toolbar>
            <TypographyTitle variant="h6">
              Remo WebUI
            </TypographyTitle>
          </Toolbar>
        </AppBar>
      </Box>


      <PaperContent>
        <Typography variant="h6" component="h2" color="primary">
          Devices
          <IconButton
            onClick={handleApplianceButtonClick}>
            <Refresh />
          </IconButton>

        </Typography>
        <AppliancesPanel appliances={appliances} />
      </PaperContent>
    </DivRoot>
  );
}