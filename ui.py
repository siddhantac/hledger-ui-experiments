import PySimpleGUI as sg
import ledger

acc = ledger.queryHledgerForAccountList("")

lb=sg.Listbox(values=acc, size=(40, 20))
sg.Window(title="Hello World", layout=[[lb]], margins=(101, 50)).read()
