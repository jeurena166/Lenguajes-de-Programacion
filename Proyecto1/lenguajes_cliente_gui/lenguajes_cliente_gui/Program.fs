open System
open System.Net
open System.Net.Sockets
open System.Text
open System.Threading
open System.Windows.Forms

let serverIpAddress = "127.0.0.1" // Replace with the actual server IP address
let serverPort = 9988 // Replace with the actual server port

let createForm () =
    let form = new Form(Visible = true, Text = "Socket Client")
    let comboBox = new ComboBox(Visible = true, Dock = DockStyle.Top)
    let textBox = new TextBox(Visible = true, Multiline = true, Dock = DockStyle.Fill, ReadOnly = true)
    form.Controls.Add(comboBox)
    form.Controls.Add(textBox)

    let receiveData () =
        try
            let client = new TcpClient()
            client.Connect(IPAddress.Parse(serverIpAddress), serverPort)
            let stream = client.GetStream()
            let buffer = Array.zeroCreate 1024
            let mutable receivedData = ""
            let mutable options : string list = [] // List to store selectable options

            while true do
                let bytesRead = stream.Read(buffer, 0, buffer.Length)
                let data = Encoding.ASCII.GetString(buffer, 0, bytesRead)
                receivedData <- receivedData + data

                // Extract options from received data (assuming each option is on a separate line)
                let newOptions = data.Split([|'\n'; '\r'|], StringSplitOptions.RemoveEmptyEntries) |> Array.toList
                options <- options @ newOptions

                // Update the GUI textbox with the receivedData on the UI thread
                form.Invoke((Action (fun () -> textBox.AppendText(data))))

                // Update the ComboBox with the new options on the UI thread
                form.Invoke((Action (fun () -> comboBox.Items.Clear())))
                form.Invoke((Action (fun () -> comboBox.Items.AddRange(options |> List.map box |> List.toArray))))


                // For debugging purposes
                printfn "Received: %s" data

                Thread.Sleep(100) // Adjust the sleep duration as needed
        with
        | ex -> printfn "Error: %s" ex.Message

    let guiThread = new Thread(new ThreadStart(receiveData))
    guiThread.Start()
    Application.Run(form)

[<STAThread>]
do
    Application.EnableVisualStyles()
    Application.SetCompatibleTextRenderingDefault(false)
    createForm()
