import { Card, Grid, Typography, Button as MuiButton, CardContent, CardHeader, Chip, Paper, styled } from "@mui/material";
import {
    Power, WindPower, Add,
    Remove, VolumeOff, EnergySavingsLeaf,
    Timer, Replay, SvgIconComponent, QuestionMark, ArrowDownward, PowerSettingsNew, ArrowUpward, Lightbulb, PowerOff, Star, Looks3, Looks4, Looks5, Looks6, LooksOne, LooksTwo
} from '@mui/icons-material';
import React, { ReactNode, useState } from "react";
import { sendSignal } from "./remo/client";
import { Appliance, Button, Signal } from "./remo/model";
import toast, { Toaster } from 'react-hot-toast';
import { stringify } from "querystring";

const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(1),
    textAlign: 'center',
    color: theme.palette.text.secondary,
}));

type SignalIconProps = {
    name: string;
    useDefaultIcon: boolean;
}

const signalIconMap = new Map<string, JSX.Element>([
    ['ico_io', <PowerSettingsNew />],
    ['ico_plus', <Add />],
    ['ico_minus', <Remove />],
    ['ico_mute', <VolumeOff />],
    ['ico_ac_eco', <EnergySavingsLeaf />],
    ['ico_return', <Replay />],
    ['ico_timer', <Timer />],
    ['ico_ac_fan', <WindPower />],
    ['ico_light', <Lightbulb />],
    ['ico_arrow_top', <ArrowUpward />],
    ['ico_arrow_bottom', <ArrowDownward />],
    ['ico_light_favorite', <Star />],
    ['ico_on', <Power />],
    ['ico_off', <PowerOff />],
    ['ico_number_1', <LooksOne />],
    ['ico_number_2', <LooksTwo />],
    ['ico_number_3', <Looks3 />],
    ['ico_number_4', <Looks4 />],
    ['ico_number_5', <Looks5 />],
    ['ico_number_6', <Looks6 />],
]);

function SignalIcon(props: SignalIconProps) {
    const icon = signalIconMap.get(props.name);
    if (!icon) {
        return props.useDefaultIcon ? (<><QuestionMark /></>) : (<></>);
    }
    return (<>{icon}</>);
}

type AppliancesPanelProps = {
    appliances: Appliance[];
}

export const AppliancesPanel: React.FC<AppliancesPanelProps> = ({
    appliances = []
}) => {
    return (
        <Grid container spacing={4} xs={12}>
            {appliances.map(
                (appliance) => {
                    if (appliance.signals == undefined || appliance.signals.length == 0) {
                        return;
                    }
                    return (<AppliancePanel appliance={appliance} />);
                }
            )}
        </Grid>
    );
};

type AppliancePanelProps = {
    appliance: Appliance;
}

const AppliancePanel: React.FC<AppliancePanelProps> = ({
    appliance,
}) => {
    if (appliance === undefined || appliance.signals === undefined) {
        return (<div>no appliances</div>)
    }
    const signals: Signal[] = (appliance.signals as Signal[])
    const buttons = appliance.signals.map((signal) => {
        return (<SignalButton name={signal.name as string} signalId={signal.id as string} icon={signal.image as string} />);
    }
    )
    return (
        <Grid item xs={4}>
            <Item>
                <Grid item xs={12} sx={{ gridArea: 'header' }}>{appliance.nickname}</Grid>
                <Grid container>
                    {signals.map(
                        (signal) => {
                            const name: string = signal.name != undefined ? signal.name as string : ""
                            const signalId: string = signal.id as string
                            const icon = signal.image as string

                            if (name == "" || signalId == "") {
                                return
                            }

                            return (
                                <Grid item xs={4}>
                                    <SignalButton name={name} signalId={signalId} icon={icon} />
                                </Grid>
                            )
                        }
                    )}
                </Grid>
            </Item>
        </Grid>
    );
};


type LightPowerStatusProps = {
    status: boolean;
}

const LightPowerStatus: React.FC<LightPowerStatusProps> = ({
    status
}) => {
    if (status) {
        return (<Chip icon={<Power />} label="ON" color="primary" />);
    }
    return (<Chip icon={<PowerOff />} label="OFF" />);
};

type SignalButtonProps = {
    name: string;
    signalId: string;
    icon: string;
};

const SignalButton: React.FC<SignalButtonProps> = ({
    name, signalId, icon
}) => {
    return (
        <>
            <Toaster
                position="top-center"
                toastOptions={{
                    duration: 5000,
                }}
            />
            <MuiButton
                startIcon={<SignalIcon name={icon} useDefaultIcon={false} />}
                onClick={() => {
                    toast.promise(
                        sendSignal(signalId),
                        { 
                            loading: 'Sending a signal...',
                            success: `Succeeded to send the signal of \`${name}\`.`,
                            error: (err) => {
                                let msg = "";
                                msg += `Failed to send the signal of \`${name}\``;
                                    msg += err !== undefined ? `\nDetail: ${err}` : '';
                                    msg += err.message !== undefined ? `\nMessage: ${err.message}` : '';
                                return msg;
                            }
                        }
                    )
                }} 
        >
            {name}
        </MuiButton>
        </>
    )
};

type LightButtonProps = {
    name: string;
    applianceId: string;
    button: string;
    icon: string;
};