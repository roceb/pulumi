// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Mypkg.Inputs
{

    /// <summary>
    /// Bastion Shareable Link.
    /// </summary>
    public sealed class BastionShareableLink : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Reference of the virtual machine resource.
        /// </summary>
        [Input("vm", required: true)]
        public string Vm { get; set; } = null!;

        public BastionShareableLink()
        {
        }
        public static new BastionShareableLink Empty => new BastionShareableLink();
    }
}
