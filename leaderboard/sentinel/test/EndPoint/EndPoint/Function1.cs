
using System.IO;
using Microsoft.AspNetCore.Mvc;
using Microsoft.Azure.WebJobs;
using Microsoft.Azure.WebJobs.Extensions.Http;
using Microsoft.AspNetCore.Http;
using Microsoft.Azure.WebJobs.Host;
using Newtonsoft.Json;
using System;

namespace EndPoint
{
    public static class Function1
    {
        private static System.Random r = new System.Random(1000); // Seed is the same

        // if you pass the provability (%), it returns false according to the probability randomly
        private static Boolean sortingHat(int probability)
        {

            var value = r.Next(100);
            return value < probability;
        }

        private static IActionResult GetResult(string label, HttpRequest req, TraceWriter log)
        {
            log.Info($"{label}: health check");
            return sortingHat(95) ? (ActionResult)new OkObjectResult("Alive")
                : new BadRequestObjectResult("Dead");
        }
      
        [FunctionName("Team")]
        public static IActionResult Run([HttpTrigger(AuthorizationLevel.Anonymous, "get", "post", Route = "team01/health")]HttpRequest req, TraceWriter log)
        {
            return GetResult("Team01", req, log);
        }
    }
}
