package no.sysco.whosnext;

import java.util.Map;
import java.util.Random;
import javax.ws.rs.Consumes;
import javax.ws.rs.POST;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.RequestBody;

/**
 *
 * @author 100tsa
 */
@Component
@Path(value = "/")
public class RestController {
    private static final Random RANDOM = new Random(System.nanoTime());
    
    @POST
    @Path("/whosnext")
    @Produces(value = "application/json")
    @Consumes(value = "application/json")
    public final Map<String, String> processRequest(@RequestBody final Map<String, String> ... payload){
        Map<String, String> nextPerson = payload[RANDOM.nextInt(payload.length)];
        System.out.println("Magic 8-Ball says:" + nextPerson);
        return nextPerson;
    }
}
