// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {pomodoro} from '../models';

export function GetPomodoroButtons():Promise<Array<pomodoro.ButtonInfo>>;

export function GetPomodoroState():Promise<pomodoro.PomodoroState>;

export function GetTimeLeft():Promise<string>;

export function PausePomodoro():Promise<void>;

export function ResumePomodoro():Promise<void>;

export function SetPomodoroTime(arg1:number):Promise<void>;

export function StartPomodoro():Promise<void>;

export function StopPomodoro():Promise<void>;
