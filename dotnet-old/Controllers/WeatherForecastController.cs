using Microsoft.AspNetCore.Mvc;

namespace dotnet_old.Controllers;

[ApiController]
[Route("/")]
public class WeatherForecastController : ControllerBase
{

  public WeatherForecastController()
  {
  }

  [HttpGet]
  public HelloResponse Get()
  {
    return new HelloResponse("world");
  }
}
