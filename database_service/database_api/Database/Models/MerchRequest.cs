using System.ComponentModel.DataAnnotations.Schema;
using System.ComponentModel.DataAnnotations;

namespace DatabaseAPI.Database.Models;

public class MerchRequest
{
	[Key]
	[DatabaseGenerated(DatabaseGeneratedOption.Identity)]
	public int Id { get; set; }

	public int UserId { get; set; }

	public string UserMessage { get; set; }

	public int? HrId { get; set; }

	public string? HrMessage { get; set; }

	public int MerchId { get; set; }

	public string StatusKey { get; set; }

	public string? ImageUrl { get; set; }


	[ForeignKey(nameof(UserId))]
	public User? User { get; set; }

	[ForeignKey(nameof(HrId))]
	public User? Hr { get; set; }

	[ForeignKey(nameof(MerchId))]
	public Merch? Merch { get; set; }

	[ForeignKey(nameof(StatusKey))]
	public Status? Status { get; set; }
}
